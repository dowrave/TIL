- 리액트 라우터에서, 동적 라우트의 동적 세그먼트에 해당하는 부분을 추출한다.
- 예를 들어 `App.tsx`에서 
```tsx
<Route path='/work' element={<WorkLayout />}>
  <Route path='study' element={<WorkStudy />} />
  <Route path='study/:id' element={<PostDetail />} />
</Route>
```
> 이렇게 구성된 라우터가 있다고 하자. 

- 이 때, `PostDetail.tsx`의 `useParams`부분은
```tsx
const { id } = useParams<{ id: number }>(); 
```
> 이렇다고 하면, `useParams`은 `App.tsx`에 있는 `:id` 를 추적, 그 곳에 해당하는 값을 추출하게 된다. 타입스크립트의 경우에만 자료형 지정`id: number`이 가능하다.


