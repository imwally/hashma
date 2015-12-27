# Hash Match

hashma is a tiny utility that helps with checking the authenticity of
files downloaded from the internet. The usual routine when downloading
an OS image from the internet is to also download another file that
contains a list of the checksums. Running a hashing algorithm against
the image file will return a string of characters, or a checksum,
which should match the same checksum found in the list.

## Usage

hashma expects two arguments, a file to verify and a file that
contains a list of checksums. It will check the MD5, SHA1, SHA256, and
SHA512 hash of the file. If the hash is found in the list of checksums
it will return the algorithm and the hash of the file, otherwise
nothing.

Please note that this utility does not verify the file containing the
checksums. It is your duty to check the authenticity of the file
containing the checksum.


```
$ hashma debian-8.2.0-i386-netinst.iso SHA256SUMS
SHA256: e4e8964aaf13a137e87de950f1ffd142ca3980a2db683534018e1ca18993be37
```



