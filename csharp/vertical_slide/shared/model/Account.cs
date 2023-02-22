namespace ex.vsa.model;

using scyna;
//using System.ComponentModel.DataAnnotations;

public class Account
{
    public ulong ID { get; set; }
    public string? Name { get; set; }
    public string? Email { get; set; }
    public string? Password { get; set; }
}