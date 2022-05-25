#!/bin/bash

printf "预备~"
sleep 1
printf "3, "
sleep 1
printf "2, "
sleep 1
printf "1"
sleep 1
printf " 开始！\n"

holdStat=true

#for i in {0..$(expr 20 \* 175 - 1)}
for i in {0..3499}
#for i in {0..174}
do
	if [[ $(expr $i % 100) -eq 0 && $i -ne 0 ]]
	then
		case $holdStat in
			true)
				holdStat=false
				;;
			false)
				holdStat=true
				;;
		esac
	fi
	case $holdStat in
		true)
			printf "\r[%d] 收！" $(expr $(expr 20 \* 180 - $i) / 20)
			;;
		false)
			printf "\r[%d] 放！" $(expr $(expr 20 \* 180 - $i) / 20)
			;;
	esac
	if [ $i -gt $(expr 20 \* 100) ]
	then
		printf "加油，胜利就在前方！"
	fi

	sleep 0.05
done

sleep 1
printf "\r恭喜你完成训练！\n"
sleep 1
printf "在本次训练中你一共获得了%d点经验！再接再厉吧！\n" $(expr $RANDOM + 23452)
sleep 1

