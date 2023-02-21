namespace ex.vsa.CreateAccount;

using scyna;

public class Handler : Endpoint.Handler<proto.CreateAccountRequest>
{
    public override void Execute()
    {
        context.Info("Receive CreateAccountRequest");
    }
}