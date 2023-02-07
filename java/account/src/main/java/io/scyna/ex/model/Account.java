package io.scyna.ex.account.model;

import io.scyna.Context;

public class Account {
    private Context context;

    public Account(Context context) {
        this.context = context;
    }

    public long ID;
    public Name name;
    public EmailAddress email;
    public Password password;
}
