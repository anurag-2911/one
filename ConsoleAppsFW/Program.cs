using System;
using System.Xml;

namespace ConsoleAppsFW
{
    class Program
    {
        const string AzureUserSourceName = "azure-logon";
        static void Main(string[] args)
        {
            string[] mechanism = new string[] { "[azure-logon:595db192-3fc9-40f2-9038-8dad125b55cd]" };
            if (mechanism != null && mechanism.Length > 0)
            {
                string azureMech = mechanism[0];
                string[] srcNtenantID = azureMech.Split(':');
                if (srcNtenantID != null && srcNtenantID.Length == 2)
                {
                    string usource = srcNtenantID[0];
                    string tenantId = srcNtenantID[1];
                    if (!string.IsNullOrEmpty(usource))
                    {
                        usource = usource.Trim(new char[] { '[', ']' });
                        tenantId = tenantId.Trim(new char[] { '[', ']' });
                    }
                }


            }



        }
    }
}
