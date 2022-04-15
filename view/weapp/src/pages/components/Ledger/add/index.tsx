import React, { useState } from "react";
import { View, Text } from "@tarojs/components"
import { AtFab, AtFloatLayout, AtTag } from "taro-ui";
import style from "./index.module.scss"
import { useReady } from "@tarojs/taro";
import IconFont from "@/utils/iconfont";
import Popup from "../../Popup";
import LedgerSelect from "../select";

type StateForm = {
  ledgerIds: number[],
  amount: number
}
const LedgerAdd = () => {

  useReady(() => {
  })

  const [state, setState] = useState<StateForm>({
    ledgerIds: [],
    amount:120
  })
  // 选择账本
  const setLedgerIds = (ledgerIds) => {
    setState({ ...state, ledgerIds })
  }

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
        onClose={() => setOpenForm(false)}
        title={<LedgerSelect selecteds={state.ledgerIds} setSelecteds={setLedgerIds} />}
        catchMove={true}
      >
        <View className={style.container}>

          <View className="flex jc-space-between ai-center">
            <View>
              <AtTag className="mr-20" active={true} type='primary'>支出</AtTag>
              <AtTag className="mr-20" active={true} type='primary'>入账</AtTag>
              <AtTag className="mr-20" active={true} type='primary'>不计入收支</AtTag>
            </View>
            <View>4月22日</View>
          </View>

          {/* 金额 */}
          <View className={style.amountBox}>
            <View className={style.unit}>￥</View>
            <View>{state.amount}<Text className={style.focus}>|</Text></View>
          </View>

          {/* 键盘 */}
        </View>



      </Popup>
      <View>

      </View>
    </>
  )
}

export default LedgerAdd
