using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class EnemyManager : MonoBehaviour
{
   
    float currentTime; // 현재 시간
    public float createTime = 1; // 일정 시간
    public GameObject enemyFactory; // 적 공장

    float minTime = 1;
    float maxTime = 5;

    // Start is called before the first frame update
    void Start()
    {
        // 태어날 때 적의 생성 시간 설정
        createTime = Random.Range(minTime, maxTime);
    }

    // Update is called once per frame
    void Update()
    {
        // 1. 시간의 흐름
        currentTime += Time.deltaTime;


        // 현재 시간이 일정 시간을 초과하면
        if (currentTime > createTime)
        {
            // 적 공장에서 적을 생성해서
            GameObject enemy = Instantiate(enemyFactory);

            // 내 위치에 갖다 놓겠다
            enemy.transform.position = transform.position;

            currentTime = 0;
            createTime = Random.Range(minTime, maxTime);

        }
    }
}
