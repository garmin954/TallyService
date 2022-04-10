import React, { useMemo, useState } from "react";
import { View, Image, Text, ScrollView, ScrollViewProps, Icon, Block } from "@tarojs/components";
import "./index.scss"
import { AtIcon } from "taro-ui";

// ScrollViewProps
const Index = (props: ScrollViewProps & { children: JSX.Element[] }) => {

  props.refresherEnabled = true;
  props.refresherThreshold = 0
  props.refresherDefaultStyle = "white"
  props.refresherBackground = "#1c9048"
  // 设置当前下拉刷新状态，true 表示下拉刷新已经被触发，false 表示下拉刷新未被触发
  const [reTriggered, setReTriggered] = useState(false)
  props.pagingEnabled = true
  props.showScrollbar = false

  const [loadStatus, setLoadStatus] = useState<'loading' | 'empty' | 'noMore' | 'wait'>('wait')
  const LoadStatusView = useMemo(() => {

    switch (loadStatus) {
      case "wait":
        return (
          <Block>
            <AtIcon className="loading" value='loading' size='16' color='#333333' />
            加载更多
          </Block>
        )
        break
      case "loading":
        return (
          <Block>
            <AtIcon className="loading" value='loading' size='16' color='#333333' />
            加载更多
          </Block>
        )
        break
      case "empty":
        return (
          <Block>

          </Block>
        )
        break
      case "noMore":
        return (
          <Block>
            我是有底线的
          </Block>
        )
        break
    }
  }, [loadStatus])

  // 下拉加载
  props.onRefresherRefresh = (event) => {
    setReTriggered(true)
    console.log(event)
    setTimeout(() => {
      setReTriggered(false)
    }, 1000)
  }

  // 上拉加载
  props.onScrollToLower = (event) => {
    console.log(event)
    setLoadStatus('loading')
    setTimeout(()=>{

      setLoadStatus('noMore')
    }, 2000)
  }

  return (
    <ScrollView
      className="gm-scroll"
      scrollY
      scrollWithAnimation
      {...props}
      refresherTriggered={reTriggered}
    >
      {props.children}

      <View className="load-more">
        {LoadStatusView}
      </View>
    </ScrollView>
  );
};

export default Index;
