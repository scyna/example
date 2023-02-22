namespace ex.vsa.repository;

using scyna;
using Scylla.Net;
using ex.vsa.model;

public class BaseRepository
{
    protected const string ACCOUNT_TABLE = "account";
    protected const string KEYSPACE = "ex_account";
    private Context context;

    protected BaseRepository(Context context) { this.context = context; }

    public static BaseRepository load(Context context)
    {
        return new BaseRepository(context);
    }

    public Account GetAccountByEmail(string email)
    {
        var query = Engine.DB.Session.Prepare(string.Format("SELECT id,email,name FROM {0}.{1} WHERE email=? LIMIT 1", KEYSPACE, ACCOUNT_TABLE));
        var statement = query.Bind(email);
        return queryAccount(statement);
    }

    public Account GetAccountByID(ulong ID)
    {
        var query = Engine.DB.Session.Prepare(string.Format("SELECT id,email,name FROM {0}.{1} WHERE id=? LIMIT 1", KEYSPACE, ACCOUNT_TABLE));
        var statement = query.Bind((long)ID);
        return queryAccount(statement);
    }

    private Account queryAccount(BoundStatement statement)
    {
        try
        {
            var rs = Engine.DB.Session.Execute(statement);
            var row = rs.First();
            var account = new Account
            {
                ID = (ulong)row.GetValue<long>("id"),
                Email = row.GetValue<string>("email"),
                Name = row.GetValue<string>("name")
            };
            return account;
        }
        catch (InvalidOperationException) { throw ex.vsa.model.Error.ACCOUNT_NOT_FOUND; }
        catch { throw scyna.Error.SERVER_ERROR; }
    }
}