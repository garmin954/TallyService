import React, { useEffect } from "react";
import { View, Button, Image } from "@tarojs/components";

import './index.scss'
import Taro, { useReady } from "@tarojs/taro";
import api from "@/api";
import { useRequest } from "taro-hooks";
import LedgerAdd from "../components/LedgerAdd";


const Index = () => {

  const { run: wxLogin, data: userInfo, loading: loadingUserInfo } = useRequest(api.wxLogin, { manual: true })
  const { run: wxRegister } = useRequest(api.wxRegister, { manual: true })

  const isLogin = () =>{
    Taro.login({
      success({ code }) {
        wxLogin({ code }).then(response => {

        })
        setTimeout(() => {
          console.log("userInfo=============================", userInfo)
        }, 2000);
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
              console.log(response)
            })
          }
        })
      }
    })
  }


  return (
    <View className="wrapper">
      {
        !userInfo?.isRegister ? (
          <Button className="button" onClick={getUserInfo}>
            login
          </Button>
        ) :null
      }

      {!loadingUserInfo ? <View>
        <Image src={userInfo?.userInfo?.avatar}></Image>
        {userInfo?.userInfo?.nickname}
        </View>:<View>loading</View>}

      <LedgerAdd />

    </View>
  );
};

export default Index;
