using System.Collections;
using System.Collections.Generic;
using JetBrains.Annotations;
using UnityEngine;

public class PlayerFire : MonoBehaviour
{
    public GameObject bulletFactory;

    // 총알 갯수
    public int poolSize = 10;

    // 배열 -> 리스트로 오브젝트 풀 변경
    public List<GameObject> bulletObjectPool;

    public GameObject firePosition;

    // 총알 오브젝트 풀 만들기
    // 1. 탄창은 시작할 때 만든다
    void Start()
    {
        // 2. 크기는 총알이 담기는 만큼
        bulletObjectPool = new List<GameObject>();
        // 3. 탄창에 넣을 개수만큼 반복
        for (int i = 0; i < poolSize; i++)
        {
            // 4. 총알 생성
            GameObject bullet = Instantiate(bulletFactory);

            // 5. 총알을 오브젝트 풀에 넣는다
            bulletObjectPool.Add(bullet);

            // 비활성화
            bullet.SetActive(false);

        }

        #if UNITY_ANDROID
                GameObject.Find("Joystick canvas XYBZ").SetActive(true);
        #elif UNITY_EDITOR || UNITY_STANDALONE
                GameObject.Find("Joystick canvas XYBZ").SetActive(false);
        #endif

    }

    // Update is called once per frame
    void Update()
    {
#if UNITY_EDITOR || UNITY_STANDALONE
        if (Input.GetButtonDown("Fire1"))
        {
            Fire();
        }
#endif
    }

    public void Fire()
    { 
        if (bulletObjectPool.Count > 0)
        {
            GameObject bullet = bulletObjectPool[0];
            bullet.SetActive(true);
            bulletObjectPool.Remove(bullet);
            // 총알 위치시키기
            bullet.transform.position = transform.position;
        }
    }
}
