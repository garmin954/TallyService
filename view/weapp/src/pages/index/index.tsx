import React from "react";
import { View, Text, Block } from "@tarojs/components";
import './index.scss'
import { AtIcon } from "taro-ui";
import { getNavBarInfo } from "../../utils/index";
import LedgerCard from "../components/LedgerCard/index"
import MScroll from "../components/Scroll/index"

const Index = () => {

  const {navBarHeight=0} = getNavBarInfo()

  return (
    <View className="home">
      <View className="statistics"  style={{height: navBarHeight+40+"px"}}></View>
      <View className="filter flex jc-space-between" style={{height: navBarHeight+50+"px"}}>
          <View className="time">
            <Text>2022年04月</Text>
            <AtIcon className="ml-10" size={18} value="chevron-down"></AtIcon>
          </View>
          {/* 类型 */}
          <View className="type">
            <Text className="txt">筛选</Text>
            {/* <Icon size='20' type='success' /> */}
          </View>
        </View>
      <View className="container" style={{paddingTop: navBarHeight+20+"px"}}>
        <MScroll style={{height:"100%"}}>
          <View className="summarize m30">
            <View className="flex mb-30">
              <View className="count-item">
                <Text className="amount">￥ 1000</Text>
                <Text className="txt">总支出</Text>
              </View>
              <View className="count-item">
                <Text className="amount">￥ 15981</Text>
                <Text className="txt">总收入</Text>
              </View>
            </View>
            <View className="split mt-40"></View>
            <View className="more flex jc-space-between">
              <Text>全部账本</Text>
              <AtIcon className="ml-10" size={18} value="chevron-right"></AtIcon>
            </View>
          </View>
          <Block>
          {
            [1,1,1,1,1,1,1,1,1,1,1,1,].map((index)=><LedgerCard key={index}  />)
          }
          </Block>
        </MScroll>
      </View>
      {/* 流水 */}
    </View>
  );
};

export default Index;
