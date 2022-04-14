import React, { useEffect, useState } from "react";
import { ScrollView, View } from "@tarojs/components"
import style from "./index.module.scss"
import Taro, { Animation, nextTick, useReady } from "@tarojs/taro";
import IconFont from "@/utils/iconfont";

type Props = {
  hideHeader?: boolean // 显示隐藏header
  visible: boolean // 显示隐藏popup
  showMask?: boolean // 显示mask
  scrollY?: boolean // 上下滑动
  scrollX?: boolean // 左右滑动
  children?: JSX.Element[] | JSX.Element // 插槽
  title?: string | JSX.Element // 标题
  onClose?: () => void // 关闭
  animationTime?: number
}

const Popup = (props: Props) => {
  const { showMask = true, scrollY = true, scrollX = true, title = '', visible = false, onClose, hideHeader = true, animationTime=500 } = props
  const [originVisible, setOriginVisible] = useState(visible)

  let timer
  useEffect(() => {
    console.log("setOriginVisible", visible)
    clearInterval(timer)
    if(!visible){
      timer = setTimeout(()=>{
        setOriginVisible(visible);
      }, animationTime)
      return
    }
    setOriginVisible(visible);

  }, [visible])

  return (
    <View
      style={{zIndex: originVisible ? 80: -1}}
      className={style.drawer}
    >
      <View
      style={{ transition: `all ${animationTime/1000}s`, opacity: `${visible ? "1" : "0"}`,display: originVisible ? "unset": 'none' }}
      className={`${style.mask} ${showMask ? style.show : ''}`} onClick={onClose}></View>
      <View
        style={{ transition: `all ${animationTime/1000}s`, transform: `translate(0px, ${visible ? "0%" : "100%"})` }}
        className={style.popup}
      >
        {hideHeader ? (
          <View className={style.header}>
            <View className={style.title}>{title}</View>
            <View className={style.close} onClick={onClose}>
              <IconFont color={"#488605"} size={50} name={"closefill"}></IconFont>
            </View>
          </View>
        ) : null}
        <ScrollView
          scrollY={scrollY}
          scrollX={scrollX}
          className={style.container}
        >
          {props.children}
        </ScrollView>
      </View>
    </View>
  )
}

export default Popup
