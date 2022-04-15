import React, { useEffect, useState } from "react";
import { CustomWrapper, ScrollView, View } from "@tarojs/components"
import style from "./index.module.scss"
import IconFont from "@/utils/iconfont";

type Props = {
  hideHeader?: boolean // 显示隐藏header
  visible: boolean // 显示隐藏popup
  showMask?: boolean // 显示mask
  scrollY?: boolean // 上下滑动
  scrollX?: boolean // 左右滑动
  children?: JSX.Element[] | JSX.Element // 插槽
  title?: string | JSX.Element // 标题
  close?: JSX.Element // 关闭
  onClose?: () => void // 关闭
  animationTime?: number // 动画时间
}

const Popup = (props: Props) => {
  const {
    showMask = true,
    scrollY = true,
    scrollX = true,
    title = <></>,
    visible = false,
    onClose,
    hideHeader = true,
    animationTime = 500,
    close = <View onClick={onClose} ><IconFont color={"#488605"} size={50} name={"closefill"} /></View>
  } = props
  const [originVisible, setOriginVisible] = useState(visible)

  let timer
  useEffect(() => {
    clearInterval(timer)
    if (!visible) {
      timer = setTimeout(() => {
        setOriginVisible(visible);
      }, animationTime)
      return
    }
    setOriginVisible(visible);

  }, [visible])


  return (
    <CustomWrapper>
      <View
        style={{ zIndex: originVisible ? 80 : -1 }}
        className={style.drawer}
      >
        <View
          style={{ transition: `all ease ${animationTime / 1000}s`, opacity: `${visible ? "1" : "0"}`, display: originVisible ? "unset" : 'none' }}
          className={`${style.mask} ${showMask ? style.show : ''}`}
          onClick={onClose}></View>
        <View
          style={{ transition: `all ease ${animationTime / 1000}s`, transform: `translate(0px, ${visible ? "0%" : "100%"})` }}
          className={style.popup}
        >
          {hideHeader ? (
            <View className={style.header}>
              <View className={style.title}>{title}</View>
              <View className={style.close} >
                {close}
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
    </CustomWrapper>
  )
}

export default Popup
