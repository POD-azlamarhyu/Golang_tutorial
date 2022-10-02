echo "ファイルを機械的に生成します"

echo "パラメータ入力"

read -p "フォルダ名>> " foldername
read -p "ファイルを複数作成 yes:0 | no:1 >> " flag
read -p "ファイル名>> " filename
read -p "ファイル数>> " filenumber
read -p "拡張子>> " types

echo "$foldername , $filename , $filenumber , $types , $flag"

if [ -n "$foldername" -a ! -d $foldername ]
then
     mkdir $foldername
fi

create_file_under_dir (){
     i=1
     while [ $i -le $filenumber ]
     do
          if [ ! -e "${foldername}/${filename}0${i}.${types}" -a "$i" -lt 10 ]
          then
               touch "${foldername}/${filename}0${i}.${types}"
          elif [ ! -e "${foldername}/${filename}0${i}.${types}" -a "$i" -lt 100 ]
          then
               touch "${foldername}/${filename}${i}.${types}"
          fi
          i=`expr $i + 1`
     done
}

create_file_current (){
     j=1
     while [ $j -le $filenumber ]
     do
          if [ ! -e "${filename}0${j}.${types}" -a "$j" -lt 10 ]
          then
               touch "${filename}0${j}.${types}"
          elif [ ! -e "${filename}0${j}.${types}" -a "$j" -lt 100 ]
          then
               touch "${filename}${j}.${types}"
          fi

          j=`expr $j + 1`
     done
}

create_multi_file (){

     if [ -n "$foldername" ]
     then
          create_file_under_dir
     else
          echo "goto create current"
          create_file_current
     fi

}

create_single_file() {
     if [ -n "$foldername"]
     then
          if [ ! -e "${foldername}/${filename}.${types}" ]
          then
               touch "${foldername}/${filename}.${types}"
          fi
     else
          if [ ! -e "${filename}.${types}" ]
          then
               touch "${filename}.${types}"
          fi
     fi
}

if [ -n "$filename" -a "$flag" -eq "0" ]
then
     create_multi_file
elif [ -n "$filename" -a "$flag" -eq "1" ]
then
     create_single_file
fi
