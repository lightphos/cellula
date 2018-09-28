package com.cellula.platform;

import java.util.UUID;

public class Nucleus {

    private String name;
    private UUID uuid;
    private long time;
    private UUID source;
    private String command;

    public Nucleus() {}

    public Nucleus(String name_, UUID uuid_) {
        name = name_;
        uuid = uuid_;
    }


    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public UUID getUuid() {
        return uuid;
    }

    public void setUuid(UUID uuid) {
        this.uuid = uuid;
    }

    public long getTime() {
        return time;
    }

    public void setTime(long time) {
        this.time = time;
    }

    public UUID getSource() {
        return source;
    }

    public void setSource(UUID source) {
        this.source = source;
    }

    public String getCommand() {
        return command;
    }

    public void setCommand(String command) {
        this.command = command;
    }

}
