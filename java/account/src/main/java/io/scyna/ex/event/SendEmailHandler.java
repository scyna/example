package io.scyna.ex.account.event;

import io.scyna.Task;
import io.scyna.ex.account.proto.SendEmailTask;

public class SendEmailHandler extends Task.Handler<SendEmailTask> {

    @Override
    public void execute() {
        context.info("Receive SendEmailTask");
        System.out.println(data);
        /* TODO */
    }
}