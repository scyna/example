package io.scyna.ex.account.domain;

import io.scyna.Context;
import io.scyna.ex.account.model.EmailAddress;
import io.scyna.ex.account.repository.Repository;
import io.scyna.ex.account.model.Error;

public class AccountService {

    public static IRepository loadRepository(Context context) {
        return Repository.load(context);
    }

    public static void assureAccountNotExists(IRepository repository, EmailAddress email) throws io.scyna.Error {
        try {
            repository.getAccountByEmail(email);
            throw Error.ACCOUNT_EXISTS;
        } catch (io.scyna.Error e) {
            if (e != Error.ACCOUNT_NOT_FOUND) {
                throw e;
            }
        }
    }
}
