# domainfinder

This tool connect to a WHOIS server to ask for details about the domains, and so, to find a domain available based on the synonyms of a certain word.

1. Build the tool executing ./build.sh
2. Define the environment variable *BIG_HUGE_THESAURUS* with the API key of the Big Huge Thesaurus api. [Get API Key here](https://words.bighugelabs.com/getkey.php)
3. Execute the binary.
4. Enter the name to search.
5. The tool shows domains that are synonyms of the entered word and its availability.

```sh
ᐅ ./domainfinder
Enter the word that represent your domain:
food
gtnutrient.net ✔
solid-foodsiite.com ✔
food-for-thoughthq.com ✔
getinteellectual-nourishment.net ✔
getcognitive-content.net ✔
contenthq.net ✖
mattertime.net ✔
mentaal-objecthq.net ✔
solid.net ✖
substancetimee.net ✔

```

