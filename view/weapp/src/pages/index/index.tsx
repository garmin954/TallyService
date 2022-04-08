import React from "react";
import { View, Image, Text, Icon } from "@tarojs/components";
import './index.scss'

const Index = () => {


  return (
    <View>
      <View className="statistics">
        <View className="filter flex">
          {/* 类型 */}
          <View className="type">
            <Text>筛选</Text>
            |
            <Icon size='20' type='success' />
          </View>
          <View className="flex">
            <View className="time">本周</View>
            <View className="time">本月</View>
            <View className="time">今年</View>
          </View>
        </View>

        <View>
          <View>总支出￥100</View>
          <View>总收入￥100</View>
        </View>
      </View>

      {/* 流水 */}
    </View>
  );
};

export default Index;
