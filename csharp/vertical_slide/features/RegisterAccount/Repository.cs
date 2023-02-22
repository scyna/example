namespace ex.vsa.CreateAccount;

using scyna;
using ex.vsa.model;
using ex.vsa.repository;

public class Repository : BaseRepository
{
    private Repository(Context context) : base(context) { }
    public static Repository create(Context context) { return new Repository(context); }

    public void AssureAccountNotExists(string email)
    {
        try
        {
            this.GetAccountByEmail(email);
        }
        catch (scyna.Error e)
        {
            if (e != model.Error.ACCOUNT_NOT_FOUND) throw;
        }
    }

    public void RegisterAccount(Account account)
    {
        var query = Engine.DB.Session.Prepare(string.Format("INSERT INTO {0}.{1}(id,email,name,password) VALUES(?,?,?,?)", KEYSPACE, ACCOUNT_TABLE));
        var statement = query.Bind((long)account.ID, account.Email, account.Name, account.Password);
        Engine.DB.Session.Execute(statement);
    }
}