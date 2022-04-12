import React from "react";
import { View, Button } from "@tarojs/components";

import './index.scss'
import Taro from "@tarojs/taro";
import api from "@/api";


const Index = () => {


  const getUserInfo = () =>{
    Taro.getUserProfile({
      desc: '用于完善会员资料', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
      success: (res) => {
        console.log(res)

        Taro.login({
          success({code}){
            api.wxLogin({code}).then(r=>{
              console.log(r)
            })
          }
        })
      }
    })
  }


  return (
    <View className="wrapper">
      <Button className="button" onClick={getUserInfo}>
        user
      </Button>
    </View>
  );
};

export default Index;
