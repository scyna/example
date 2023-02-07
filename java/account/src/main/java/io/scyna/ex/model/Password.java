package io.scyna.ex.account.model;

public class Password {
    String value;

    private Password(String password) {
        this.value = password;
    }

    public static Password create(String password) throws io.scyna.Error {
        if (password == null || password.length() == 0) {
            throw Error.BAD_PASSWORD;
        }

        /* TODO: some other rules here */

        return new Password(password);
    }

    public String toString() {
        return value;
    }
}
