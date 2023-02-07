package io.scyna.ex.account.domain;

import io.scyna.Command;
import io.scyna.ex.account.model.Account;
import io.scyna.ex.account.model.EmailAddress;

public interface IRepository {

    void createAccount(Command command, Account account) throws io.scyna.Error;

    Account getAccountByEmail(EmailAddress email) throws io.scyna.Error;

    Account getAccountByID(long ID) throws io.scyna.Error;

}
