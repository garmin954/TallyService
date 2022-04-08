import React from "react";
import { View, Image, Text, Icon } from "@tarojs/components";
import './index.scss'
import { AtIcon } from "taro-ui";
import { getNavBarInfo } from "../../utils/index";

const Index = () => {

  const {navBarHeight=0} = getNavBarInfo()

  return (
    <View>
      <View className="statistics" style={{paddingTop: navBarHeight+10+"px"}}>
        <View className="filter flex jc-space-between">
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
      </View>

      {/* 流水 */}
    </View>
  );
};

export default Index;
