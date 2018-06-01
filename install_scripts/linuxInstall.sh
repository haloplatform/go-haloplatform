#!/bin/bash
echo ""
echo ""
echo "Welcome To:"
echo ""
echo ""
echo "      ___           ___       ___     "
echo "     /  /\         /  /\     /  /\    "
echo "    /  /::\       /  /:/    /  /::\   "
echo "   /  /:/\:\     /  /:/    /  /:/\:\  "
echo "  /  /:/  \:\   /  /:/    /  /:/  \:\ "
echo " /__/:/_\_ \:\ /__/:/    /__/:/ \__\:\\"
echo " \  \:\__/\_\/ \  \:\    \  \:\ /  /:/"
echo "  \  \:\ \:\    \  \:\    \  \:\  /:/ "
echo "   \  \:\/:/     \  \:\    \  \:\/:/  "
echo "    \  \::/       \  \:\    \  \::/   "
echo "     \__\/         \__\/     \__\/    "
echo ""
echo ""
echo "Go Halo Platform - Binary Install - Linux"

sleep 3

echo ""
echo ""
echo "############################################################"
echo "###### Step 1: Downloading Glo Binary from Website    ######"
echo "############################################################"
curl https://binaries.haloplatform.tech/glo/latest-linux-x64.tar.gz --output latest-linux-x64-temp.tar.gz

sleep 1 

echo ""
echo ""
echo "############################################################"
echo "###### Step 2: Untar Glo Binary                       ######"
echo "############################################################"
tar -zxvf latest-linux-x64-temp.tar.gz

sleep 1 

echo ""
echo ""
echo "############################################################"
echo "###### Step 3: Moving glo to /usr/local/bin           ######"
echo "############################################################"
mv glo-*-linux-x64 /usr/local/bin/glo
chmod +x /usr/local/bin/glo

sleep 1 

echo ""
echo ""
echo "############################################################"
echo "###### Step 4: Deleting Temp glo binary               ######"
echo "############################################################"
rm latest-linux-x64-temp.tar.gz

sleep 1

echo ""
echo ""                                                   
echo " # #    #  ####  #####   ##   #      #      ###### #####  "
echo " # ##   # #        #    #  #  #      #      #      #    # "
echo " # # #  #  ####    #   #    # #      #      #####  #    # "
echo " # #  # #      #   #   ###### #      #      #      #    # "
echo " # #   ## #    #   #   #    # #      #      #      #    # "
echo " # #    #  ####    #   #    # ###### ###### ###### #####  "
echo ""
echo " #########################################################"
echo ""
echo "Instructions:"
echo ""
echo "To use 'glo' please simply type 'glo --datadir <give_it_a_name>'."
echo ""
echo "To automatically attach your console type 'glo --datadir <give_it_a_name> console'."
echo "Glo will automatically connect to the console for you. See the repository for more info."
echo ""
echo "Github Repo: https://github.com/haloplatform/go-haloplatform"
                                                          

