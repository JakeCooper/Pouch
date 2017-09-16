#! /bin/bash
APP="Pouch"
EXT="pouch"
COMMENT="$APP's data file"

# Create mime xml (this doesnt work)
#echo "<?xml version=\"1.0\" encoding=\"UTF-8\"?>
#<mime-info xmlns=\"http://www.freedesktop.org/standards/shared-mime-info\">
#    <mime-type type=\"application/x-$APP\">
#        <comment>$COMMENT</comment>
#        <icon name=\"application-x-$APP\"/>
#        <glob pattern=\"*.$EXT\"/>
#    </mime-type>
#</mime-info>" > /usr/share/mime/application/x-$EXT.xml

# Create application desktop
echo "[Desktop Entry]
Name=$APP
Exec=Pouch %U
MimeType=application/x-$APP
Icon=~/.local/share/icons/hicolor/48x48/apps/$EXT.png
Terminal=false
Type=Application"> ~/.local/share/applications/$APP.desktop

chmod 711 ~/.local/share/applications/$APP.desktop

# copy associated icons to the icons folder
# This also doesnt work........
cp ./static/$EXT.png ~/.local/share/icons/hicolor/48x48/apps/

# update databases for both application and mime
update-desktop-database ~/.local/share/applications
update-mime-database    ~/.local/share/mime
