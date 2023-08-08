using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class DestroyZone : MonoBehaviour
{

    private void OnTriggerEnter(Collider other)
    {
        // 부딪힌 상대의 이름이 Bullet이라면
        if (other.gameObject.name.Contains("Bullet") ||
            other.gameObject.name.Contains("Enemy"))
        {
            // 부딪힌 물체 비활성화
            other.gameObject.SetActive(false);


            if (other.gameObject.name.Contains("Bullet"))
            {
                // PlayerFire 클래스 얻어오기
                PlayerFire player = GameObject.Find("Player").GetComponent<PlayerFire>();

                // 리스트에 총알 삽입
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
