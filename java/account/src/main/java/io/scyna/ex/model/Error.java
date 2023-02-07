package io.scyna.ex.account.model;

public class Error {
    public static final io.scyna.Error ACCOUNT_EXISTS = io.scyna.Error.New(100, "Account Exists");
    public static final io.scyna.Error ACCOUNT_NOT_FOUND = io.scyna.Error.New(101, "Account Not Found");
    public static final io.scyna.Error BAD_EMAIL = io.scyna.Error.New(102, "Bad Email");
    public static final io.scyna.Error BAD_NAME = io.scyna.Error.New(103, "Bad Name");
    public static final io.scyna.Error BAD_PASSWORD = io.scyna.Error.New(104, "Bad Password");
}
