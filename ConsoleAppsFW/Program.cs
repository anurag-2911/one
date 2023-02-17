using System;
using System.Linq;
using System.Security.Cryptography;
using System.Security.Principal;

namespace ConsoleAppsFW
{
    class Program
    {
        static byte[] s_additionalEntropy = { 9, 8, 7, 6, 5 };

        static void Main(string[] args)
        {
            string clientID = "25356a13-d1c9-43cb-858e-88fc11631b29";

            byte[] clientIDBytes = clientID.Select(b => (byte)b).ToArray();

            byte[] protectedtClientID = Protect(clientIDBytes);

            char[] clientIDcharacters = protectedtClientID.Select(b => (char)b).ToArray();
            string protectedstring = new string(clientIDcharacters);

            byte[] protectedstringbytes = protectedstring.Select(b => (byte)b).ToArray();

            byte[] uprototectedClientID = Unprotect(protectedstringbytes);

            string origClientID = new string(uprototectedClientID.Select(b => (char)b).ToArray());

            Console.ReadKey();


        }
        public static byte[] Protect(byte[] data)
        {
            try
            {
                // Encrypt the data using DataProtectionScope.CurrentUser. The result can be decrypted
                // only by the same current user.
                return ProtectedData.Protect(data, s_additionalEntropy, DataProtectionScope.CurrentUser);
            }
            catch (CryptographicException e)
            {

                return null;
            }
        }

        public static byte[] Unprotect(byte[] data)
        {
            try
            {
                //Decrypt the data using DataProtectionScope.CurrentUser.
                return ProtectedData.Unprotect(data, s_additionalEntropy, DataProtectionScope.CurrentUser);
            }
            catch (CryptographicException e)
            {

                return null;
            }
        }
    }
}
