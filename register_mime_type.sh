#! bin/bash
APP="Pouch"
EXT="pouch"
COMMENT="$APP's data file"

# Create directories if missing
mkdir -p ~/.local/share/mime/packages
mkdir -p ~/.local/share/applications

# Create mime xml 
echo "<?xml version=\"1.0\" encoding=\"UTF-8\"?>
<mime-info xmlns=\"http://www.freedesktop.org/standards/shared-mime-info\">
    <mime-type type=\"application/x-$APP\">
        <comment>$COMMENT</comment>
        <icon name=\"application-x-$APP\"/>
        <glob pattern=\"*.$EXT\"/>
    </mime-type>
</mime-info>" > ~/.local/share/mime/packages/application-x-$APP.xml

# Create application desktop
echo "[Desktop Entry]
Name=$APP
Exec=/usr/bin/$APP %U
MimeType=application/x-$APP
Icon=$APP
Terminal=false
Type=Application
Categories=
Comment=
"> ~/.local/share/applications/$APP.desktop

# update databases for both application and mime
update-desktop-database ~/.local/share/applications
update-mime-database    ~/.local/share/mime

# set default application. `mimeopen <filepath>` now opens things properly
touch test.pouch
echo "2\nPouch" | mimeopen -d test.pouch
rm test.pouch

# copy associated icons to pixmaps
cp ./static/$APP.png                ~/.local/share/pixmaps
cp ./static/application-x-$APP.png  ~/.local/share/pixmaps