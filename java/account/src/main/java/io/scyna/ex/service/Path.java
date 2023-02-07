package io.scyna.ex.account.service;

public abstract class Path {
    public static String CREATE_ACCOUNT_URL = "/account/create";
    public static String GET_ACCOUNT_BY_EMAIl_URL = "/account/get-by-email";
    public static String GET_ACCOUNT_URL = "/account/get";
    public static String ACCOUNT_CREATED_CHANNEL = "account_created";
    public static String SEND_EMAIL_CHANNEL = "send_email";
}