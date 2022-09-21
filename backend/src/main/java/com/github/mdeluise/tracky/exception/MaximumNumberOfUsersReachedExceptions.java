package com.github.mdeluise.tracky.exception;

public class MaximumNumberOfUsersReachedExceptions extends RuntimeException {
    public MaximumNumberOfUsersReachedExceptions() {
        super("Maximum number of user reached.");
    }
}
