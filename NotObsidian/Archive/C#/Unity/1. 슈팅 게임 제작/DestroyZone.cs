using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class DestroyZone : MonoBehaviour
{

    private void OnTriggerEnter(Collider other)
    {
        // �ε��� ����� �̸��� Bullet�̶��
        if (other.gameObject.name.Contains("Bullet") ||
            other.gameObject.name.Contains("Enemy"))
        {
            // �ε��� ��ü ��Ȱ��ȭ
            other.gameObject.SetActive(false);


            if (other.gameObject.name.Contains("Bullet"))
            {
                // PlayerFire Ŭ���� ������
                PlayerFire player = GameObject.Find("Player").GetComponent<PlayerFire>();

                // ����Ʈ�� �Ѿ� ����
                player.bulletObjectPool.Add(other.gameObject);

            }
            // enemy 
            else if (other.gameObject.name.Contains("Enemy"))
            {
                GameObject emObject = GameObject.Find("EnemyManager");
                EnemyManager manager = emObject.GetComponent<EnemyManager>();

                manager.enemyObjectPool.Add(other.gameObject);
            }
        }


    }
}
