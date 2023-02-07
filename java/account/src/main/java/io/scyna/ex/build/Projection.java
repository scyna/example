package io.scyna.ex.account.build;

import io.scyna.Engine;
import io.scyna.Event;
import io.scyna.Task;
import io.scyna.ex.account.event.AccountCreatedHandler;
import io.scyna.ex.account.event.SendEmailHandler;
import io.scyna.ex.account.service.Path;

public class Projection {

    public static void main(String[] args) throws Exception {
        Engine.init("http://127.0.0.1:8081", "ex_account", "123456");
        Event.register("ex_account", Path.ACCOUNT_CREATED_CHANNEL, new AccountCreatedHandler());
        Task.register("ex_account", Path.SEND_EMAIL_CHANNEL, new SendEmailHandler());
        Event.startListening();
    }
}
