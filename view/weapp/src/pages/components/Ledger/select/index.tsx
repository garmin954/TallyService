import { View, Text } from "@tarojs/components"
import React, { useEffect, useState } from "react"
import Popup from "@/pages/components/Popup"
import IconFont from "@/utils/iconfont"
import { useRequest } from "taro-hooks"
import api from "@/api"
import { useReady } from "@tarojs/taro"
import { AtCheckbox } from "taro-ui"

const LedgerSelect = () => {
  const [openForm, setOpenForm] = useState(false)

  const [ledgerOption, setLedgerOption] = useState([])
  const [selectedList, setSelectedList] = useState([])


  const { data, loading, run: fetchLedgerUser } = useRequest(api.fetchLedgerUser, { manual: true, cacheKey: `ledger_user`, formatResult:(res)=>res.data as any })
  useEffect(()=>{
    if(data){
      const options = data.map(({name,id, desc})=>({
        value:id,
        label:name,
        desc:desc,
      }))
      setLedgerOption(options)
    }
  }, [data])

  useReady(() => {
    fetchLedgerUser({})
  })
  // 打开账本
  const openLedgerChoose = () => {
    setOpenForm(true)
  }

  // 确认关闭
  const confirm = () => {
    setOpenForm(false)
  }

  const CloseSlot = () => {
    return (<View onClick={() => confirm()}>完成</View>)
  }

  // 选择
  const changeLedgerOption = (e) =>{
    setSelectedList(e)
    console.log(e)
  }

  return (<>
    <View onClick={() => openLedgerChoose()} className="flex ai-center">
      <Text className="mr-5">我的账本</Text>
      <IconFont size={25} name={"arrow-down-filling"} />
    </View>
    <Popup
      animationTime={600}
      visible={openForm}
      onClose={() => setOpenForm(false)}
      close={<CloseSlot />}
    >
      <View className="">

      {loading ? <>loading</> : <AtCheckbox
        options={ledgerOption}
        selectedList={selectedList}
        onChange={changeLedgerOption}
      />}
      </View>
    </Popup>
  </>)
}

export default LedgerSelect
