#!/bin/sh
#xdg-mime install pouch.xml --novendor
#sudo update-mime-database /usr/share/mime
cp application-pouch.xml ~/.local/share/mime/packages/
cp pouch.desktop ~/.local/share/applications/
update-mime-database ~/.local/share/mime
update-desktop-database ~/.local/share/applications
