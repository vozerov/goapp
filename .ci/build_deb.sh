#!/bin/bash
set -e

buildid=$1
dpath=web-calc
[ -e $dpath ] && rm -r $dpath
[ -f "web-calc.deb" ] && rm -f web-calc.deb

mkdir -p $dpath $dpath/DEBIAN $dpath/usr/bin $dpath/opt/goapp/ $dpath/etc/systemd/system
echo -e "Package: web-calc\nVersion: $buildid\nMaintainer: ttserg\nSection: misc\nDescription: web calc\nArchitecture: all"  > $dpath/DEBIAN/control
echo -e "#!/bin/bash\nsystemctl daemon-reload\nsystemctl start web-calc\nsystemctl enable web-calc"  > $dpath/DEBIAN/postinst

chmod 0755 $dpath/DEBIAN/postinst

cp goapp_build $dpath/opt/goapp/

FILE=$dpath/etc/systemd/system/web-calc.service
cat << EOF > $FILE
[Unit]
Description=Web Calc
After=network.target
[Service]
Type=simple
ExecStart=/opt/goapp/goapp_build
[Install]
WantedBy=multi-user.target
EOF

md5deep -r $dpath > $dpath/DEBIAN/md5sums

fakeroot dpkg-deb --build $dpath

mv $dpath.deb web-calc.deb

exit 0
