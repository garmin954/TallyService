import React, { useState } from "react";
import { View, Text } from "@tarojs/components"
import { AtFab, AtFloatLayout } from "taro-ui";
import style from "./index.module.scss"
import { useReady } from "@tarojs/taro";
import IconFont from "@/utils/iconfont";
import Popup from "../../Popup";
import LedgerSelect from "../select";

const LedgerAdd = () => {

  useReady(() => {
  })

  const [openForm, setOpenForm] = useState(false)
  return (
    <>
      <View className={style.addBtn}>
        <AtFab onClick={() => setOpenForm(true)}>
          <IconFont color="#fff" name="wirt" size={50} />
        </AtFab>
      </View>

      <Popup
        animationTime={600}
        visible={openForm}
        onClose={() => setOpenForm(false)} title={<LedgerSelect />}>
        <View style={{ height: '70vh' }}>

        </View>

      </Popup>
      <View>

      </View>
    </>
  )
}

export default LedgerAdd
