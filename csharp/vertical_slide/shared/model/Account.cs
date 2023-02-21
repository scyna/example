namespace ex.vsa.model;

using scyna;

public class Account
{
    private scyna.Context context;

    public Account(Context context)
    {
        this.context = context;
    }

    public ulong ID;
    public Name? Name;
    public EmailAddress? Email;
    public Password? Password;

}