foosync
=======

Synchronize your foobar's playback statistics based on your lastfm scrobbles

For now, it can only use lastfm json scrobbles, I intend to add support for regular .scrobbler.log

For now, it only updates the last_played info, I intend to add playcount

There might be a ±1hour difference from the real time because of DST (daylight saving time)

# Compiling
- [http://golang.org/doc/install](golang)

# Requires
- [http://www.foobar2000.org/components/view/foo_texttools](foo_texttools)
- [http://www.lastfm.fr/settings/dataexporter](lastfm scrobbles)

# Usage
1. Open foobar without any song playing
2. In foobar, select all your music files (for example with an autoplaylist of all your libraries)
3. Right click, Utilities, Text Tools, Advanced and use this pattern (you can skip duplicates)
    $if(%first_played%,$if($and(%artist%,%album%,%title%),%artist%  %album% %title% $replace(%first_played%, ,T)Z))
4. Copy and paste into a new file called ps.tsv
5. Export a full backup of your playback statistics in a file called ps.xml
6. Put your lastfm json scrobbles into a folder called scrobbles
7. Run foosync.exe
8. A file called ps_updated.xml is created, import it with playback statistics

# Files and Folders architecture recap
    - foosync.exe
    - ps.tsv
    - ps.xml
    \ scrobbles
      - scrobbles_XXXX-XX.json
      - scrobbles_XXXX-XX.json