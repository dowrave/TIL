import { ItemType } from './App2'
import logo from './logo.svg'
import { useDeferredValue } from 'react'

type props = {results : Array<ItemType>};

const ItemList = (props : props) => {
    const deferredResults = useDeferredValue(props.results);
    // const divRows = props.results.map((item) => {
    const divRows = deferredResults.map((item) => (
        <div key={item.id}>
            id: {item.id}
            <br />
            keyword: {item.keyword}
            <br />
            <img src={logo} style={{ width: 100, height: 100}} />
        </div>
    ))

    return <div>{divRows}</div>
}

export default ItemList;