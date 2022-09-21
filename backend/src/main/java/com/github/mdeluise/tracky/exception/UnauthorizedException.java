package com.github.mdeluise.tracky.exception;

public class UnauthorizedException extends RuntimeException {
    public UnauthorizedException(String message) {
        super(message);
    }

    public UnauthorizedException() {
        this("Operation not authorized");
    }
}
