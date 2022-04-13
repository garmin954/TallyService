import React, { useState } from "react";
import {View, Text} from "@tarojs/components"
import { AtFab, AtFloatLayout } from "taro-ui";
import style from "./index.module.scss"
import { useReady } from "@tarojs/taro";
import IconFont from "@/utils/iconfont";

const LedgerAdd = () =>{

  useReady(()=>{
  })

  const [openForm, setOpenForm] = useState(false)
  return (
    <>
      <View className={style.addBtn}>
        <AtFab onClick={()=>setOpenForm(true)}>
          <IconFont color="#fff"  name="wirt" size={50} />
        </AtFab>
      </View>

      <AtFloatLayout isOpened={openForm}  title="这是个标题" onClose={()=>setOpenForm(false)}>
        这是内容区 随你怎么写这是内容区 随你怎么写这是内容区 随你怎么写这是内容区
        随你怎么写这是内容区 随你怎么写这是内容区 随你怎么写
      </AtFloatLayout>
    </>
  )
}

export default LedgerAdd
