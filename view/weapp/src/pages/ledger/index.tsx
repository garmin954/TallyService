import React from "react";
import { View, Image, Text } from "@tarojs/components";
import './index.scss'
import { AtAvatar } from 'taro-ui'
const ledgerImg = require("./ledger.png")

const Index = () => {


  return (
    <View>
      <AtAvatar image='https://jdc.jd.com/img/200a'></AtAvatar>
      <View>

        {/* 时间 */}

        {/* 统计 */}

        {/* 流水 */}
        <View className="ledger-box">
          <View className="ledger">
            <Image src={ledgerImg} className="cover"></Image>
            <View className="intro clearfix">
              <Text className="name mb-20">我的账本</Text>
              <View className='at-row count-box clearfix'>
                <View className='at-col count-item'>
                  <Text className="count">10</Text>
                  <Text className="sort-t">总天数</Text>
                </View>
                <View className='at-col count-item'>
                  <Text className="count">32</Text>
                  <Text className="sort-t">总笔数</Text>
                </View>
                <View className='at-col count-item'>
                  <Text className="count">
                    690.50
                    <Text className="fontsize-10 ml-5">元</Text>
                  </Text>
                  <Text className="sort-t">总金额</Text>
                </View>
              </View>
            </View>
          </View>
          <View className="close">x</View>
        </View>
      </View>
    </View>
  );
};

export default Index;
