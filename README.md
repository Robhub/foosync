foosync
=======

Synchronize your foobar's playback statistics based on your lastfm scrobbles

# Compiling
- [http://golang.org/doc/install](golang)

# Requires
- [http://www.foobar2000.org/components/view/foo_texttools](foo_texttools)

# Usage
1. In foobar, select all your music files (for example with an autoplaylist of all your libraries)
2. Right click, Utilities, Text Tools, Advanced and use this pattern (you can skip duplicates)
    $if(%first_played%,$if($and(%artist%,%album%,%title%),%artist%  %album% %title% $replace(%first_played%, ,T)Z))
4. Copy and paste into a new file called ps.tsv
5. Export a full backup of your playback statistics in a file called ps.xml




# Files and Folders architecture recap
    - foosync.exe
    - ps.tsv
    - ps.xml
    \ scrobbles
      - scrobbles_XXXX-XX.json
      - scrobbles_XXXX-XX.json