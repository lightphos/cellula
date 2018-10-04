package com.cellula.platform.component;

import io.vertx.core.AbstractVerticle;
import io.vertx.core.Handler;
import io.vertx.core.http.HttpServer;
import io.vertx.core.http.HttpServerRequest;

public class Ping extends AbstractVerticle {

    public void start() {
// mvn archetype:generate -Dfilter=io.vertx:

        HttpServer server = vertx.createHttpServer()
                .requestHandler(new Handler<HttpServerRequest>() {
                    @Override
                    public void handle(HttpServerRequest httpServerRequest) {
                        httpServerRequest.response().end("We are cooking");
                    }
                });

        server.listen(8080, "localhost", res -> {
            if (res.succeeded()) {
                System.out.println("Server is now listening on actual port: " + server.toString());
            } else {
                System.out.println("Failed to bind!");
            }
        });

        System.out.println("Webserver started, listening on port: 8080");

    }

}
