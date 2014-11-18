## == Strings ==

addcslashes — Quote string with slashes in a C style  
addslashes — Quote string with slashes   
bin2hex — Convert binary data into hexadecimal representation   
chop — Alias of rtrim
chr — Return a specific character   
chunk_split — Split a string into smaller chunks    
convert_cyr_string — Convert from one Cyrillic character set to another    
convert_uudecode — Decode a uuencoded string    
convert_uuencode — Uuencode a string    
count_chars — Return information about characters used in a string    
crc32 — Calculates the crc32 polynomial of a string    
crypt — One-way string hashing    
~~echo — Output one or more strings~~    
explode — Split a string by string    
fprintf — Write a formatted string to a stream    
get_html_translation_table — Returns the translation table used by htmlspecialchars and htmlentities    
hebrev — Convert logical Hebrew text to visual text    
hebrevc — Convert logical Hebrew text to visual text with newline conversion    
hex2bin — Decodes a hexadecimally encoded binary string    
html_entity_decode — Convert all HTML entities to their applicable characters    
htmlentities — Convert all applicable characters to HTML entities    
htmlspecialchars_decode — Convert special HTML entities back to characters    
htmlspecialchars — Convert special characters to HTML entities    
~~implode — Join array elements with a string~~    
~~join — Alias of implode~~    
lcfirst — Make a string's first character lowercase    
levenshtein — Calculate Levenshtein distance between two strings    
localeconv — Get numeric formatting information    
~~ltrim — Strip whitespace (or other characters) from the beginning  of a string~~    
md5_file — Calculates the md5 hash of a given file    
~~md5 — Calculate the md5 hash of a string~~    
metaphone — Calculate the metaphone key of a string    
money_format — Formats a number as a currency string    
nl_langinfo — Query language and locale information    
nl2br — Inserts HTML line breaks before all newlines in a string    
number_format — Format a number with grouped thousands    
ord — Return ASCII value of character    
parse_str — Parses the string into variables    
~~print — Output a string~~    
printf — Output a formatted string    
quoted_printable_decode — Convert a quoted-printable string to an  8 bit string    
quoted_printable_encode — Convert a 8 bit string to a quoted-printable string    
quotemeta — Quote meta characters    
~~rtrim — Strip whitespace (or other characters) from the end of a string~~    
setlocale — Set locale information    
sha1_file — Calculate the sha1 hash of a file     
sha1 — Calculate the sha1 hash of a string    
similar_text — Calculate the similarity between two strings    
soundex — Calculate the soundex key of a string    
sprintf — Return a formatted string    
sscanf — Parses input from a string according to a format    
str_getcsv — Parse a CSV string into an array    
str_ireplace — Case-insensitive version of str_replace.    
str_pad — Pad a string to a certain length with another string    
str_repeat — Repeat a string    
~~str_replace — Replace all occurrences of the search string with the replacement string~~    
str_rot13 — Perform the rot13 transform on a string    
str_shuffle — Randomly shuffles a string    
str_split — Convert a string to an array    
str_word_count — Return information about words used in a string    
strcasecmp — Binary safe case-insensitive string comparison    
strchr — Alias of strstr    
strcmp — Binary safe string comparison    
strcoll — Locale based string comparison    
strcspn — Find length of initial segment not matching mask    
strip_tags — Strip HTML and PHP tags from a string    
stripcslashes — Un-quote string quoted with addcslashes    
stripos — Find the position of the first occurrence of a case-insensitive substring in a string    
stripslashes — Un-quotes a quoted string    
stristr — Case-insensitive strstr      
~~strlen — Get string length~~    
strnatcasecmp — Case insensitive string comparisons using a "natural order" algorithm    
strnatcmp — String comparisons using a "natural order" algorithm    
strncasecmp — Binary safe case-insensitive string comparison of the first n characters    
strncmp — Binary safe string comparison of the first n characters    
strpbrk — Search a string for any of a set of characters    
strpos — Find the position of the first occurrence of a substring in a string    
strrchr — Find the last occurrence of a character in a string    
strrev — Reverse a string    
strripos — Find the position of the last occurrence of a case-insensitive substring in a string    
strrpos — Find the position of the last occurrence of a substring in a string    
strspn — Finds the length of the initial segment of a string consisting entirely of characters contained within a given mask.    
strstr — Find the first occurrence of a string    
strtok — Tokenize string    
strtolower — Make a string lowercase    
strtoupper — Make a string uppercase    
strtr — Translate characters or replace substrings    
substr_compare — Binary safe comparison of two strings from an offset, up to length characters    
substr_count — Count the number of substring occurrences    
substr_replace — Replace text within a portion of a string    
~~substr — Return part of a string~~    
~~trim — Strip whitespace (or other characters) from the beginning  and end of a string~~    
ucfirst — Make a string's first character uppercase    
ucwords — Uppercase the first character of each word in a string     
vfprintf — Write a formatted string to a stream    
vprintf — Output a formatted string     
vsprintf — Return a formatted string    
wordwrap — Wraps a string to a given number of characters    

## == Date ==

checkdate — Validate a Gregorian date   
date_add — Alias of DateTime::add   
date_create_from_format — Alias of DateTime::createFromFormat   
date_create_immutable_from_format — Alias of DateTimeImmutable::createFromFormat   
date_create_immutable — Alias of DateTimeImmutable::__construct   
date_create — Alias of DateTime::__construct   
date_date_set — Alias of DateTime::setDate   
date_default_timezone_get — Gets the default timezone used by all date/time functions in a script   
date_default_timezone_set — Sets the default timezone used by all date/time functions in a script   
date_diff — Alias of DateTime::diff   
date_format — Alias of DateTime::format   
date_get_last_errors — Alias of DateTime::getLastErrors  
date_interval_create_from_date_string — Alias of DateInterval::createFromDateString   
date_interval_format — Alias of DateInterval::format   
date_isodate_set — Alias of DateTime::setISODate   
date_modify — Alias of DateTime::modify   
date_offset_get — Alias of DateTime::getOffset   
date_parse_from_format — Get info about given date formatted according to the specified format   
date_parse — Returns associative array with detailed info about given date   
date_sub — Alias of DateTime::sub   
date_sun_info — Returns an array with information about sunset/sunrise and twilight begin/end   
date_sunrise — Returns time of sunrise for a given day and location   
date_sunset — Returns time of sunset for a given day and location    
date_time_set — Alias of DateTime::setTime   
date_timestamp_get — Alias of DateTime::getTimestamp   
date_timestamp_set — Alias of DateTime::setTimestamp   
date_timezone_get — Alias of DateTime::getTimezone   
date_timezone_set — Alias of DateTime::setTimezone   
date — Format a local time/date    
getdate — Get date/time information   
gettimeofday — Get current time   
gmdate — Format a GMT/UTC date/time   
gmmktime — Get Unix timestamp for a GMT date   
gmstrftime — Format a GMT/UTC time/date according to locale settings   
idate — Format a local time/date as integer   
localtime — Get the local time   
microtime — Return current Unix timestamp with microseconds   
mktime — Get Unix timestamp for a date   
strftime — Format a local time/date according to locale settings   
strptime — Parse a time/date generated with strftime   
strtotime — Parse about any English textual datetime description into a Unix timestamp   
~~time — Return current Unix timestamp~~   
timezone_abbreviations_list — Alias of DateTimeZone::listAbbreviations   
timezone_identifiers_list — Alias of DateTimeZone::listIdentifiers   
timezone_location_get — Alias of DateTimeZone::getLocation   
timezone_name_from_abbr — Returns the timezone name from abbreviation   
timezone_name_get — Alias of DateTimeZone::getName   
timezone_offset_get — Alias of DateTimeZone::getOffset   
timezone_open — Alias of DateTimeZone::__construct   
timezone_transitions_get — Alias of DateTimeZone::getTransitions   
timezone_version_get — Gets the version of the timezonedb   







