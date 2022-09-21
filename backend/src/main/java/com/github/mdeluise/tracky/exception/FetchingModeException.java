package com.github.mdeluise.tracky.exception;

public class FetchingModeException extends RuntimeException {
    public FetchingModeException() {
        super("The used fetching mode does not provide this function");
    }
}
