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
it will return the hash and file name otherwise nothing.

Please note that this utility does not verify the file containing the
checksums. It is your duty to check the authenticity of the file
containing the checksum. Usually this is done via GPG signatures.


```
$ hashma debian-live-8.2.0-amd64-gnome-desktop.iso SHA256SUMS
aea5b49206904f3a2a01b31aee185f32fe46324b0f2bacbd259fc74374cb7b62  debian-live-8.2.0-amd64-gnome-desktop.iso
```



