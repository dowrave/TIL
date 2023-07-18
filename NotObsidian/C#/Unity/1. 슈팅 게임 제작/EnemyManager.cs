using System.Collections;
using System.Collections.Generic;
using System.Linq;
using Unity.VisualScripting;
using UnityEngine;

public class EnemyManager : MonoBehaviour
{

    // 오브젝트 풀 만들기
    public int poolSize = 10;
    public List<GameObject> enemyObjectPool;
    public Transform[] spawnPoints;

    float currentTime; // 현재 시간
    public float createTime; // 일정 시간
    public GameObject enemyFactory; // 적 공장

    public float minTime = 0.2f;
    public float maxTime = 0.5f;



    void Start()
    {
        // 태어날 때 적의 생성 시간 설정
        createTime = Random.Range(minTime, maxTime);

        // 오브젝트 풀 관련
        enemyObjectPool = new List<GameObject>();
        for (int i = 0; i < poolSize; i++)
        {
            GameObject enemy = Instantiate(enemyFactory);
            enemyObjectPool.Add(enemy);
            enemy.SetActive(false);
        }
    }

    // Update is called once per frame
    void Update()
    {
        // 1. 시간의 흐름
        currentTime += Time.deltaTime;


        // 현재 시간이 일정 시간을 초과하면
        if (currentTime > createTime)
        {
            // 오브젝트 풀 관련 반복문 : active인 것만 포지션에 옮긴다
            if (enemyObjectPool.Count > 0)
            {
                GameObject enemy = enemyObjectPool[0];
                enemyObjectPool.Remove(enemy);

                enemy.transform.position = transform.position;

                // spawnpoint 관련
                int index = Random.Range(0, spawnPoints.Length); // spawnpoint 중 1개 지정
                enemy.transform.position = spawnPoints[index].position;
                enemy.SetActive(true);

            }

            currentTime = 0;
            createTime = Random.Range(minTime, maxTime);

        }
    }
}
