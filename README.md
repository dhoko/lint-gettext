# Linter-gettext

Lint a pot file to see if it is valid or nor and print the error.

ex:
```sh
$ ./lint-gettext -input template.pot
[+] lint file template.pot
Error inside the file:
  "template.pot:4112: end-of-line within string"

Details:
     4107: msgstr ""
     4108:
     4109: #: apps/Payments/app/Services/SubscriptionService.php:615
     4110: msgctxt "Error"
     4111: msgid "Your subscription is managed by Mozilla.
>>>  4112:           Please contact Mozilla support to change your subscription"
     4113: msgstr ""
     4114:
     4115: #: apps/Payments/app/Services/SubscriptionService.php:672
     4116: msgctxt "Error"
     4117: msgid "Currency mismatch"
```

API:

```sh
lint-gettext -input <filePath>
```


Dependencies:
- [gettext](https://www.gnu.org/software/gettext/)

Debian/Ubuntu
```sh
apt install gettext
```

Mac
```sh
brew install gettext
brew link --force gettext
```
