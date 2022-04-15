import React, { useEffect } from "react";
import { View, Button, Image } from "@tarojs/components";

import './index.scss'
import Taro, { useReady } from "@tarojs/taro";
import api from "@/api";
import { useRequest } from "taro-hooks";
import LedgerAdd from "../components/Ledger/add";
import { useState } from "react";


const Index = () => {

  const [userInfo, setUserInfo] = useState<UserInfo>()
  const { run: wxLogin, loading: loadingUserInfo } = useRequest(api.wxLogin, { manual: true })
  const { run: wxRegister } = useRequest(api.wxRegister, { manual: true })

  const isLogin = () =>{
    Taro.login({
      success({ code }) {
        wxLogin({ code }).then(response => {
          console.log("response", response)
          if(response.data.isRegister){
            setUserInfo(response.data.userInfo)
          }
        })
      }
    })
  }

  useEffect(()=>{
    console.log("userInfo-------------------", userInfo)
  }, [userInfo])

  useReady(()=>{
    isLogin()
  })

  const getUserInfo = () => {
    Taro.getUserProfile({
      desc: '用于完善会员资料', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
      success: ({ encryptedData, iv }) => {
        Taro.login({
          success({ code }) {
            wxRegister({ code, encryptedData, iv }).then(response => {
              setUserInfo(response.data)
            })
          }
        })
      }
    })
  }


  return (
    <View className="wrapper">
      {
        !userInfo ? (
          <Button className="button" onClick={getUserInfo}>
            login
          </Button>
        ) :null
      }

      {!loadingUserInfo && userInfo ? <View>
        <Image src={userInfo?.avatar}></Image>
        {userInfo.nickname}
        </View>:<View>loading</View>}

      <LedgerAdd />

    </View>
  );
};

export default Index;
