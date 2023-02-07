package io.scyna.ex.account.build;

import io.scyna.Endpoint;
import io.scyna.Engine;
import io.scyna.ex.account.service.GetAccountByEmailService;
import io.scyna.ex.account.service.Path;

public class Reader {
    public static void main(String[] args) throws Exception {
        Engine.init("http://127.0.0.1:8081", "ex_account", "123456");
        Endpoint.register(Path.GET_ACCOUNT_BY_EMAIl_URL, new GetAccountByEmailService());
    }
}
