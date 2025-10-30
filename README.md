How to run my code in Lab_week03
- I was writing the MD5, SHA1, SHA512 under the same utils/crack and by using the same wordlist file.
  
- How to run the MD5
  cd cd "C:\Users\DELL\OneDrive - Cambodia Academy of Digital Technology\Desktop\Go\Cryptography_Lab_Week3\lab1"
  go run . -mode=md5 -target="6a85dfd77d9cb35770c9dc6728d73d3f" -wordlist="./wordlist.txt" -verbose="verbose_md5.txt" 2>&1 | Tee-Object -FilePath verbose_md5.txt
  
- How to run the SHA
  cd cd "C:\Users\DELL\OneDrive - Cambodia Academy of Digital Technology\Desktop\Go\Cryptography_Lab_Week3\lab1"
  go run main.go -mode=sha1 -target=aa1c7d931cf140bb35a5a16adeb83a551649c3b9 -wordlist="wordlist.txt" 2>&1 | Tee-Object -FilePath verbose2.txt
  
- How to run the SHA512
  cd cd "C:\Users\DELL\OneDrive - Cambodia Academy of Digital Technology\Desktop\Go\Cryptography_Lab_Week3\lab1"
  go run main.go -mode=sha512 -target="<hash>" -wordlist="wordlist.txt" -verbose="verbose_sha512.txt"
