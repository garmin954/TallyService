import React from "react";
import { View, Text } from "@tarojs/components";
import "./index.scss"
import IconFont from "@/utils/iconfont";



const Index = () => {

  return (
    <View className="ledger-card mt-20 m30">
      <View className="flex jc-space-between">
        <View className="time">04月25日 星期二</View>
        <View className="statistics-ledger">支 28.99 收 0.00</View>
      </View>
      <View className="record mt-20">
        <View className="record-item flex">
          <IconFont name="bonus" size={50} />
          <View className="intro ml-10">
            <View className="header flex jc-space-between">
              <Text className="title">沃尔玛超市消费</Text>
              <Text className="amount">￥ 99</Text>
            </View>
            <View className="tag-box">
              生活日用
            </View>
            <View className="certificate"></View>
          </View>
        </View>
      </View>
    </View>
  );
};

export default Index;
