package com.cellula.platform;

import io.vertx.core.*;
import io.vertx.core.eventbus.EventBus;
import io.vertx.core.impl.Args;
import io.vertx.core.json.Json;
import io.vertx.core.json.JsonObject;

import io.vertx.core.spi.cluster.ClusterManager;
import io.vertx.spi.cluster.hazelcast.HazelcastClusterManager;

import java.util.ArrayList;
import java.util.UUID;
import java.util.logging.Level;
import java.util.logging.Logger;

public class Cell extends AbstractVerticle {

    private Logger logger = Logger.getLogger(Cell.class.getName());
    private static final String CQUE = "nucleus.list";
    private UUID myID;


    private void log(String s) {
        logger.log(Level.INFO, myID +": " +s);
    }


    ArrayList<String> dlist = new ArrayList<>();
    ArrayList<Nucleus> deployers = new ArrayList<>();

    public void start() {

        String cells = config().getString("cells");
        for (String v : cells.split(" ")) dlist.add(v);
        deploy(dlist);
    }


    int countDz;
    int countRx;
    Nucleus selected = null;
    public Nucleus nucleus(String d) {
        Nucleus c = new Nucleus();
        c.setUuid(UUID.randomUUID());
        c.setName(d);
        c.setSource(myID);
        c.setTime(System.nanoTime());
        log(c.getName() + ":" + c.getUuid() + " -=========> Distribute Time is " + c.getTime());
        return c;
    }


    // master requests services of the rest of the cells
    public void serviceRequestForRemaining(EventBus eb) {
        int ix = 1;
        for (Nucleus c : deployers) {
            c.setName(dlist.get(ix++));
            c.setCommand("serve");
//            log("Go serve " + c.getSource() + " " + c.getName());
            pub(eb, c);
        }
    }

    public void pub(EventBus eb, Nucleus c) {
        JsonObject jo = new JsonObject(Json.encode(c));
        eb.publish(CQUE, jo);
    }


    boolean deployed = false;
    private void deploySelf(Vertx vx, Nucleus c) {
        if (!deployed) { // ensure only once
            deployed = true;
            vx.deployVerticle(c.getName());
            log("Deployed: running " + c.getName());
        }
    }


    long timerID=0;
    public void deploy(ArrayList dlist) {
        ClusterManager mgr = new HazelcastClusterManager();
        VertxOptions options = new VertxOptions().setClusterManager(mgr);
        Vertx.clusteredVertx(options, res -> {
            if (res.succeeded()) {

                Vertx vx = res.result();
                EventBus eb = vertx.eventBus();
                log("------------------- We now have a clustered event bus: " + eb);


                eb.consumer(CQUE, message -> {

                    JsonObject svert = (JsonObject) message.body();

        //            if (svert.equals("*")) {
        //                log("Rx: I have run " + selected.getName() + " " + selected.getTimestamp());
        //                return;
        //            }


                    //log("Rx: " + svert.toString());
                    Nucleus vert = Json.decodeValue(svert.toString(), Nucleus.class);

                    // got a request from the master to serve the rest of the set
                    if ("serve".equals(vert.getCommand())) {
                        vertx.cancelTimer(timerID); // cancel retry as we have a master
                        // is the source id ours -> start up our nucleus, each nucleus starts their own nucleus. the master
                        // has already started its
                        if (myID.equals(vert.getSource())) {
                            // one for us
                            selected = vert;
                            deploySelf(vx, selected);
                        }

                        return;
                    }




                    if (vert.getSource().compareTo(selected.getSource()) == 0) {
                        // ignore our messages
                        return;
                    }


                    // is the vertical the one I've selected
                    if (vert.getName().equals(selected.getName())) {
        //                log("Add " + vert.getSource() + " " + vert.getName());
                        deployers.add(vert); // keep a list of cells
                        countRx++;

                        if (countRx == (countDz - 1)) { // we have the first nucleus request from all the cells
                            // check which was the quickest
                            Nucleus quickest = selected;
                            for (Nucleus d : deployers) {
                                if (d.getTime() < quickest.getTime()) quickest = d;
                            }

                            // was this one the quickest - if so we are the boss
                            if (quickest.getSource() == selected.getSource()) {
                                // we are master
                                log("CountRx " + countRx + " ***master*** deploy  " + selected.getName());
                                vertx.cancelTimer(timerID); // stop retrying
                                deploySelf(vx, selected);

                                // kick off servant requests --> tell the others to run the rest of the cells
                                serviceRequestForRemaining(eb);

                            }
                        }
                    }
                    else { // keep a list of
                        log("Got outbound request for nucleus " + vert.getName());
                    }


                });

                myID = UUID.randomUUID();
                countDz = dlist.size();
                countRx = 0;
                selected = nucleus((String) dlist.get(0)); // create a nucleus from the first one

                log("Publish first");
                pub(eb, selected); // send it to the Q, first one

                //wait 10 secs. and send again
                timerID = vertx.setPeriodic(10000, new Handler<Long>() {

                    @Override
                    public void handle(Long id) {
                        log("Timer 1 fired: " + id);
                        pub(eb, selected); // first one
                    }
                });

            } else {
                System.out.println("Failed: " + res.cause());
            }
        });


//        eb.publish(CQUE, "*");

//        for (String d : dlist) {
//            vx.deployVerticle(d, res -> {
//                if (res.succeeded()) {
//                    log("Deployment id is: " + res.result());
//                } else {
//                    log("Deployment failed!");
//                }
//            });
//        }

    }


    public static void main(String[] args) {
// http://atomix.io/copycat/
// use counter to check contention for leader election.
//
        // running a jar file
//        // ProcessBuilder pb = new ProcessBuilder("/path/to/java", "-jar", "your.jar", "arg1", "arg2");
//        pb.directory(new File("preferred/working/directory"));
//        Process p = pb.start();
        //
        // automatically deploy to docker
        // https://github.com/docker-java/docker-java/wiki
        String commandLine = "run com.cellula.platform.Cell -conf src/main/conf/deployer-conf.json -instances 4";
        String[] sargs = commandLine.split(" ");
        Args nargs = new Args(sargs);
        Launcher.main(sargs);

    }
}
