using System.Collections;
using System.Collections.Generic;
using System.Linq;
using Unity.VisualScripting;
using UnityEngine;

public class EnemyManager : MonoBehaviour
{

    // ������Ʈ Ǯ �����
    public int poolSize = 10;
    public List<GameObject> enemyObjectPool;
    public Transform[] spawnPoints;

    float currentTime; // ���� �ð�
    public float createTime; // ���� �ð�
    public GameObject enemyFactory; // �� ����

    public float minTime = 0.2f;
    public float maxTime = 0.5f;



    void Start()
    {
        // �¾ �� ���� ���� �ð� ����
        createTime = Random.Range(minTime, maxTime);

        // ������Ʈ Ǯ ����
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
        // 1. �ð��� �帧
        currentTime += Time.deltaTime;


        // ���� �ð��� ���� �ð��� �ʰ��ϸ�
        if (currentTime > createTime)
        {
            // ������Ʈ Ǯ ���� �ݺ��� : active�� �͸� �����ǿ� �ű��
            if (enemyObjectPool.Count > 0)
            {
                GameObject enemy = enemyObjectPool[0];
                enemyObjectPool.Remove(enemy);

                enemy.transform.position = transform.position;

                // spawnpoint ����
                int index = Random.Range(0, spawnPoints.Length); // spawnpoint �� 1�� ����
                enemy.transform.position = spawnPoints[index].position;
                enemy.SetActive(true);

            }

            currentTime = 0;
            createTime = Random.Range(minTime, maxTime);

        }
    }
}
