using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class EnemyManager : MonoBehaviour
{
   
    float currentTime; // ���� �ð�
    public float createTime = 1; // ���� �ð�
    public GameObject enemyFactory; // �� ����

    float minTime = 1;
    float maxTime = 5;

    // Start is called before the first frame update
    void Start()
    {
        // �¾ �� ���� ���� �ð� ����
        createTime = Random.Range(minTime, maxTime);
    }

    // Update is called once per frame
    void Update()
    {
        // 1. �ð��� �帧
        currentTime += Time.deltaTime;


        // ���� �ð��� ���� �ð��� �ʰ��ϸ�
        if (currentTime > createTime)
        {
            // �� ���忡�� ���� �����ؼ�
            GameObject enemy = Instantiate(enemyFactory);

            // �� ��ġ�� ���� ���ڴ�
            enemy.transform.position = transform.position;

            currentTime = 0;
            createTime = Random.Range(minTime, maxTime);

        }
    }
}
